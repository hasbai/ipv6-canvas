<template>
  <main id="main">
    <PixelInfo :pixel="pixel" :offset="offset" :ratio="ratio" :active="pixelInfoActive"></PixelInfo>
    <HomeLoader v-if="imgLoading"></HomeLoader>
    <canvas
        id="canvas"
        v-else
        :width="canvasSize.x" :height="canvasSize.y"
        @wheel="onWheel"
        @mousemove="onMouseMove"
        @mousedown="onMouseDown"
        @mouseup="onMouseUp"
    >
    </canvas>
    <BottomPanel :current="current" :image-size="imgSize" :ratio="ratio"></BottomPanel>
  </main>
</template>


<script setup lang="ts">
import BottomPanel from "@/components/BottomPanel.vue"
import PixelInfo from "@/components/PixelInfo.vue"
import HomeLoader from "@/components/HomeLoader.vue"
import {useMessage} from "naive-ui"
import {Point} from "@/models/point"
import {nextTick, onMounted, reactive, ref} from "vue"
import {isValidCoordinate} from "@/utils"
import {Color, Pixel} from "@/models/pixel"

function updateCoordinate(e: MouseEvent) {
  current.set(new Point(e.offsetX, e.offsetY).sub(offset).div(ratio.value).floor())
}

function drawPixel(pixel: Pixel) {
  const d = fixedCtx.createImageData(1, 1)
  d.data[0] = pixel.color.r
  d.data[1] = pixel.color.g
  d.data[2] = pixel.color.b
  d.data[3] = pixel.color.a
  fixedCtx.putImageData(d, pixel.coordinate.x, pixel.coordinate.y)
  drawImage()
}

function drawImage(r: number = ratio.value, p: Point = current) {
  if (r !== ratio.value) {
    offset.sub(p.copy().times(r - ratio.value))
    ratio.value = r
  }
  ctx.imageSmoothingEnabled = false
  clearImage()
  ctx.drawImage(
      fixedCanvas,
      offset.x,
      offset.y,
      Math.floor(imgSize.x * r),
      Math.floor(imgSize.y * r),
  )
}

function clearImage() {
  ctx.fillStyle = 'grey'  // clear canvas
  ctx.fillRect(0, 0, canvasSize.x, canvasSize.y)
}

function moveImage(p: Point) {
  offset.add(p)
  drawImage()
}

let mouseDown = false
let mouseLocation = new Point(0, 0)

async function onKeyDown(e: KeyboardEvent) {
  const _move = (px = 4) => {
    if (e.key === 'w' || e.key === 'ArrowUp') {
      moveImage(new Point(0, px))
    } else if (e.key === 'a' || e.key === 'ArrowLeft') {
      moveImage(new Point(px, 0))
    } else if (e.key === 's' || e.key === 'ArrowDown') {
      moveImage(new Point(0, -px))
    } else if (e.key === 'd' || e.key === 'ArrowRight') {
      moveImage(new Point(-px, 0))
    }
  }
  for (let i = 0; i < 12; i++) {
    _move()
    await new Promise((r) => setTimeout(r, 2))
  }
  hidePixelInfo()
}

async function onClick(e: MouseEvent) {
  updateCoordinate(e)
  hidePixelInfo()
  if (!isValidCoordinate(current, imgSize)) {
    return
  }
  // transition
  let r = ratio.value
  const step = r < targetRatio ? 0.2 : -0.2
  while (r > targetRatio + 0.01 || r < targetRatio - 0.01) {
    r += step
    drawImage(r, current)
    await new Promise((r) => setTimeout(r, 1))
  }
  drawImage(targetRatio, current)
  showPixelInfo()
}

function onMouseMove(e: MouseEvent) {
  if (mouseDown && isValidCoordinate(current, imgSize)) {
    moveImage(new Point(e.movementX, e.movementY))
    hidePixelInfo()
  }
  updateCoordinate(e)
}

function onMouseDown(e: MouseEvent) {
  mouseDown = true
  mouseLocation = new Point(e.clientX, e.clientY)
}

function onMouseUp(e: MouseEvent) {
  mouseDown = false
  if (mouseLocation.equals(new Point(e.clientX, e.clientY))) {
    onClick(e)
  }
}

async function onWheel(e: WheelEvent) {
  let level = -e.deltaY > 0 ? 1 : -1 // 正值为放大，负值为缩小
  level *= Math.ceil(ratio.value / 8)
  drawImage(Math.max(ratio.value + level, 1), current)
}

const pixel = reactive(new Pixel(new Point(0, 0), new Color(0, 0, 0, 0)))
const pixelInfoActive = ref(false)

function showPixelInfo() {
  pixel.coordinate = current.copy()
  pixel.color.fromCanvas(fixedCtx, current)
  pixelInfoActive.value = true
}

function hidePixelInfo() {
  pixelInfoActive.value = false
}

function connectWs() {
  const ws = new WebSocket(
      window.location.origin.replace("http", "ws") + "/ws"
  );
  ws.binaryType = "arraybuffer";
  ws.onopen = () => {
    console.log("websocket connected");
  };
  ws.onerror = () => {
    console.log("websocket connect failed");
  };
  ws.onclose = () => {
    console.log("websocket disconnected");
  };
  ws.onmessage = onMessage;
}

function onMessage(e: MessageEvent) {
  const data = new Uint8Array(e.data);
  switch (data[0]) {
    case MessageType.Pixel:
      const pixel = Pixel.fromBytes(data);
      console.log(pixel)
      drawPixel(pixel);
      break;
  }
}

const MessageType = {
  Pixel: 0x1,
  MetaData: 0x2,
};

const targetRatio = 24
const ratio = ref(1)
const imgLoading = ref(true)
const current = reactive(new Point(0, 0))
const offset = reactive(new Point(0, 0))
const canvasSize = reactive(new Point(0, 0))
const imgSize = reactive(new Point(0, 0))
let ctx: CanvasRenderingContext2D
let fixedCanvas: HTMLCanvasElement
let fixedCtx: CanvasRenderingContext2D

// window.$message = useMessage()
connectWs()
document.addEventListener('keydown', onKeyDown)

onMounted(async () => {
  console.log('mounted')
  // load image
  let img = new Image()
  img.src = '/image'
  img.crossOrigin = 'anonymous'
  await new Promise((r) => (img.onload = r))
  imgSize.set(new Point(img.naturalWidth, img.naturalHeight))
  imgLoading.value = false
  await nextTick()

// init canvas
  const main = document.getElementById('main')!
  canvasSize.set(new Point(main.offsetWidth, main.offsetHeight))
  ctx = (document.getElementById('canvas') as HTMLCanvasElement).getContext('2d')!
  clearImage()
  await nextTick()

// create a fixed canvas
  fixedCanvas = document.createElement('canvas')
  fixedCanvas.width = imgSize.x
  fixedCanvas.height = imgSize.y
  fixedCtx = fixedCanvas.getContext('2d', {willReadFrequently: true})!
  fixedCtx.drawImage(img, 0, 0, imgSize.x, imgSize.y)

// draw to the real canvas
  offset.set(canvasSize.copy().sub(imgSize).div(2).floor())
  drawImage()
})

</script>

<style scoped>
</style>
