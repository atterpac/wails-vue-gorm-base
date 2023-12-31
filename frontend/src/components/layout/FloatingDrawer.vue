<template>
    <section class="floating-drawer" :style='{display: props.visible ? "flex":"none", top: props.top, width: props.width }'>
        <HueShift />
    </section>
</template>

<script setup lang='ts'>
import HueShift from '../HueShift.vue'

const props = defineProps({
    top: {
        type: String,
        default: '5%'
    },
    right: {
        type: String,
        default: '16px'
    },
    left: {
        type: String,
        default: 'auto'
    },
    width: {
        type: String,
        default: '200px'
    },
    height: {
        type: String,
        default: '90%'
    },
    visible: Boolean
    })
</script>

<style scoped>
    .floating-drawer {
        width: 100px;
        height: 25%;
        z-index: 100;
        position: fixed;
        top: 5%;
        right: 16px;
        border-radius: 8px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

        border: 0.5px solid var(--primary-dark);

        /* Box Shadow */
        box-shadow: 0px 0px 20px 0px hsl(var(--hue), 100%, 50%);
        -webkit-box-shadow: 0px 0px 15px 0px hsla(var(--hue), 0%, 20%, 0.75);
        background-color: var(--background-color-light);
    }
</style>
