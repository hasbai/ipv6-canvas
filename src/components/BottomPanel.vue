<template>
  <div class="bottom">
    <div class="info">
      <div class="align">
        <Icon size="1rem">
          <LocationOnFilled/>
        </Icon>
        {{ coordinateStr }}
      </div>
      <div class="align online">
        <Icon size="1rem">
          <PeopleAltFilled/>
        </Icon>
      </div>
    </div>
    <ColorSelector @updatePixel="onUpdatePixel"></ColorSelector>
  </div>
</template>

<script setup lang="ts">
  import ColorSelector from "@/components/ColorSelector.vue"
  import {LocationOnFilled, PeopleAltFilled} from '@vicons/material'
  import {Icon} from '@vicons/utils'
  import {Point} from "@/models/point"
  import {isValidCoordinate} from "@/utils"
  import {computed} from "vue"

  const props = defineProps({
    current: {
      required: true,
      type: Point,
    },
    imageSize: {
      required: true,
      type: Point,
    },
    ratio: {
      required: true,
      type: Number,
    },
  })

  const coordinateStr = computed(() => {
    if (isValidCoordinate(props.current, props.imageSize)) {
      return props.current.toString()
    } else {
      return '(-, -)'
    }
  })

  function onUpdatePixel(color: string) {
    console.log(color)
  }

</script>

<style scoped>
    .bottom {
        display: flex;
        flex-direction: column;
        align-items: center;
        position: fixed;
        bottom: 0;
    }

    .info {
        display: flex;
        align-items: center;
        width: max-content;
        padding: 0.25rem 0.5rem;
        border-radius: 0.25rem 0.25rem 0 0;
        background-image: linear-gradient(-180deg, rgba(241, 241, 241, 0.8) 0%, rgba(227, 227, 227, 0.8) 100%);
    }

    .align {
        display: flex;
        align-items: center;
    }

    .online {
        margin-left: 1rem;
    }
</style>
