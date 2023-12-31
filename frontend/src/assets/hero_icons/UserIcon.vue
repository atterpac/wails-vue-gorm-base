<template>
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" :fill="props.hero_fill" :stroke="props.hero_stroke" :width="props.hero_width" :height="props.hero_height">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z" />
        </svg>
    </div>
</template>

<script lang="ts" setup>
const props = defineProps({
    hero_stroke: {
        type: String,
        default: '#000000'
    },
    hero_fill: {
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
