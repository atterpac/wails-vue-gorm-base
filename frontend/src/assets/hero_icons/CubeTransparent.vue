<template>
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" :stroke="props.hero_stroke" :width="props.hero_width" :height="props.hero_height">
            <path stroke-linecap="round" stroke-linejoin="round" d="m21 7.5-2.25-1.313M21 7.5v2.25m0-2.25-2.25 1.313M3 7.5l2.25-1.313M3 7.5l2.25 1.313M3 7.5v2.25m9 3 2.25-1.313M12 12.75l-2.25-1.313M12 12.75V15m0 6.75 2.25-1.313M12 21.75V19.5m0 2.25-2.25-1.313m0-16.875L12 2.25l2.25 1.313M21 14.25v2.25l-2.25 1.313m-13.5 0L3 16.5v-2.25" />
        </svg>
    </div>
</template>

<script lang="ts" setup>
const props = defineProps({
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
