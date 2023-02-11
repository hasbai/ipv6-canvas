<template>
  <div
      class="pixel-info"
      v-if="active"
      :style="{
        height: ratio + 'px',
        width: ratio + 'px',
        top: frameTop + 'px',
        left: frameLeft + 'px'
      }"
  >
    <PixelData :props="props.pixel" :is-loading="isLoading"></PixelData>
  </div>
</template>

<script setup lang="ts">
import PixelData from "@/components/PixelData.vue"
import {computed} from "vue"
import {Pixel} from "@/models/pixel"
import {Point} from "@/models/point"

const props = defineProps({
  pixel: {type: Pixel, required: true},
  ratio: {type: Number, required: true},
  offset: {type: Point, required: true},
  active: {type: Boolean, default: false},
})

const frameLeft = computed(() => {
  return props.offset.x + props.pixel.coordinate.x * props.ratio - 1
})
const frameTop = computed(() => {
  return props.offset.y + props.pixel.coordinate.y * props.ratio - 1
})

let isLoading = false

</script>

<style scoped>
.pixel-info {
  position: absolute;
  border: 1px solid black;
}


</style>
