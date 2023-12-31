<template>
  <div>
    <label>Hue:</label>
    <input v-model="hue" type="range" min="0" max="360" :style="{ background: `linear-gradient(to right, hsl(0, ${saturation}% , ${lightness}%), hsl(60, ${saturation}%, ${lightness}%), hsl(120, ${saturation}%, ${lightness}%), hsl(180, ${saturation}%, ${lightness}%), hsl(240, ${saturation}%, ${lightness}%), hsl(300, ${saturation}%, ${lightness}%), hsl(360, ${saturation}%, ${lightness}%))` }">

    <label>Saturation:</label>
    <input v-model="saturation" type="range" min="0" max="100">

    <label>Lightness:</label>
    <input v-model="lightness" type="range" min="0" max="100">
  </div>
</template>

<script setup lang='ts'>
import { ref, watch } from 'vue';

const hue = ref(150);
const saturation = ref(50);
const lightness = ref(50);

const updateRootStyles = () => {
  document.documentElement.style.setProperty('--hue', hue.value.toString());
  document.documentElement.style.setProperty('--saturation', `${saturation.value}%`);
  document.documentElement.style.setProperty('--lightness', `${lightness.value}%`);
};

watch([hue, saturation, lightness], () => {
  updateRootStyles();
});
</script>

<style scoped>

/* Add any additional styles for your component here */
label {
  display: block;
  margin-bottom: 0.5rem;
}
</style>

