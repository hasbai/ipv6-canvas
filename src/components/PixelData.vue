<template>
  <div class="pixel">
    <div class="coordinates bold">
      <span>X：{{ props.coordinate.x }}</span>
      <span>Y：{{ props.coordinate.y }}</span>
    </div>
    <hr>
    <div class="color bold">
      <span>HEX</span>
      <div class="value">
        <span style="margin-right: 0.25em">{{ props.color.toString() }}</span>
        <span class="color-block" :style="{backgroundColor: props.color.toString()}">向晚</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

import type {Pixel} from "@/models/pixel";

const {props} = defineProps<{ props: Pixel }>()

</script>

<style scoped>
.pixel {
  position: relative;
  left: 125%;
  top: -50%;
  width: max-content;
  background-color: rgba(204, 204, 204, 0.5);
  padding: 1em;
}

/* 此时一个点大小为 24 px */
.pixel:before {
  content: "";
  width: 0;
  height: 0;
  border-top: 6px solid transparent;
  border-bottom: 6px solid transparent;
  border-right: 6px solid rgba(204, 204, 204, 0.5);
  position: absolute;
  left: -6px;
}

.coordinates {
  display: flex;
  justify-content: space-around;
}

.color {
  display: flex;
  justify-content: space-between;
}

.color .value {
  margin-left: 1em;
}


.color-block {
  color: transparent;
  margin-left: 0.25em;
}

hr {
  border: 0.5px solid rgba(128, 128, 128, 0.42);
}

.spinner {
  width: 4rem;
  height: 4rem;
  margin: 1.5rem;
  position: relative;
}

.double-bounce1, .double-bounce2 {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background-color: #333;
  opacity: 0.6;
  position: absolute;
  top: 0;
  left: 0;

  -webkit-animation: sk-bounce 2.0s infinite ease-in-out;
  animation: sk-bounce 2.0s infinite ease-in-out;
}

.double-bounce2 {
  -webkit-animation-delay: -1.0s;
  animation-delay: -1.0s;
}

@keyframes sk-bounce {
  0%, 100% {
    transform: scale(0.0);
  }
  50% {
    transform: scale(1.0);
  }
}
</style>
