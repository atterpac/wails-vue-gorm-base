<template>
    <section class="floating-drawer" :style='{display: props.visible ? "flex":"none" }'>
        <h3> I am a floating drawer </h3>
    </section>
</template>

<script setup lang='ts'>

const props = defineProps({
    visible: Boolean
    })
</script>

<style scoped>
    .floating-drawer {
        width: 300px;
        height: 90%;
        z-index: 100;
        position: fixed;
        top: 5%;
        right: 48px;
        border-radius: 8px;

        background-color: darkcyan;
    }
</style>
