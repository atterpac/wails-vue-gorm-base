<template>
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" :stroke="props.hero_stroke" :fill='props.hero_fill' :width="props.hero_width" :height="props.hero_height">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4.098 19.902a3.75 3.75 0 0 0 5.304 0l6.401-6.402M6.75 21A3.75 3.75 0 0 1 3 17.25V4.125C3 3.504 3.504 3 4.125 3h5.25c.621 0 1.125.504 1.125 1.125v4.072M6.75 21a3.75 3.75 0 0 0 3.75-3.75V8.197M6.75 21h13.125c.621 0 1.125-.504 1.125-1.125v-5.25c0-.621-.504-1.125-1.125-1.125h-4.072M10.5 8.197l2.88-2.88c.438-.439 1.15-.439 1.59 0l3.712 3.713c.44.44.44 1.152 0 1.59l-2.879 2.88M6.75 17.25h.008v.008H6.75v-.008Z" />
        </svg>
    </div>
</template>

<script lang="ts" setup>
const props = defineProps({
    hero_fill: {
        type: String,
        default: '#000000'
    },
    hero_stroke: {
        type: String,
        default: '#000000'
    },
    hero_width: {
        type: String,
        default: '24px'
    },
    hero_height: {
        type: String,
        default: '24px'
    }
})
</script>

<style scoped>
</style>
