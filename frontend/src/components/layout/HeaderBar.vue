<template>
    <section class="header" :style='{height: props.height}' >
    </section>
</template>

<script setup lang='ts'>
    import ArrowDown from '../../assets/hero_icons/ArrowDown.vue'
    const props = defineProps({
        height: String 
    })

    console.log(props.height)
</script>

<style scoped>
    .header {
        width: 100%;
        position: fixed;
        top: 0;
        left: 0;
        z-index: 100;
        align-items: center;
        justify-content: center;
        margin: auto;
    }
</style>
