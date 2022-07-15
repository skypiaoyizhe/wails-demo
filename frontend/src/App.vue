<script setup>
import HelloWorld from './components/HelloWorld.vue'
import {Test} from '../wailsjs/go/main/App'
const handleTest = () => {
  console.log("点击测试")
  Test().then(res => {
    console.log("结果",res)
  })
}
</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <button @click="handleTest">点击测试</button>
  <HelloWorld/>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
