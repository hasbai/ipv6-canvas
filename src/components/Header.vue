<template>
  <header>
    <div class="github">
      <n-button
          href="https://github.com/hasbai/ipv6-canvas"
          tag="a"
          target="view_window"
          :bordered="false"
      >
        <span style="margin-right: 0.5rem">Github Repo</span>
        <n-icon :size="'1.5rem'">
          <GithubFilled></GithubFilled>
        </n-icon>
      </n-button>
    </div>
  </header>
</template>

<script setup lang="ts">
import {NButton, NIcon} from 'naive-ui'
import {GithubFilled} from '@vicons/antd'

</script>

<style scoped></style>
