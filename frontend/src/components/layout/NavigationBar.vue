<template>
    <section class="nav-bar" :style="{ width: props.isCollapse ? '24px' : '180px' }">
        <CubeTransparent class="logo" hero_stroke="var(--primary-color)" hero_width='48px' hero_height='48px'/>
        <div class='nav-item'>
            <button 
                class='nav-button'
                v-for="item in navigationItems" 
                :key=item.path
                @click='onClick(item.path)'>
                <div class='icon'>
                    <Component :is='item.icon' hero_stroke="var(--primary-color)" hero_width='24px' hero_height='24px' />
                </div>
                <div class='name'>
                    {{item.name}}
                </div>
            </button>
        </div>
    </section>
</template>

<script lang="ts" setup>
import type { RouteLocationRaw } from 'vue-router';
import { router } from '../../router';
import CubeTransparent from '../../assets/hero_icons/CubeTransparent.vue'
import HomeIcon from '../../assets/hero_icons/HomeIcon.vue'
import UserIcon from '../../assets/hero_icons/UserIcon.vue'
import ColorSwatch from '../../assets/hero_icons/ColorSwatch.vue'
import type { Component } from 'vue';

type NavigationItem = {
    name: string;
    path: string;
    icon?: Component
}

const navigationItems: NavigationItem[] = [
    {
        name: 'Home',
        path: '/',
        icon: HomeIcon
    },
    {
        name: 'Auth',
        path: '/login',
        icon: UserIcon
    },
    {
        name: 'Colors',
        path: '/colors',
        icon: ColorSwatch
    },
]

const props = defineProps({
    isCollapse: {
        type: Boolean,
        default: false
    }
})

const onClick = (path: RouteLocationRaw) => {
    router.push(path);
}
</script>

<style scoped>
.nav-bar {
    position: fixed;
    left: 0;
    top: 24px;
    align-items: center;
    justify-content: center;
    height: 100%;
    z-index: 100;
}

.nav-item {
    height: 36px;
}

.nav-button {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    height: 100%;
}

.icon {
    padding-right: 10px;
}

button {
    width: 100%;
    height: 36px;
    border: none;
    background-color: transparent;
    color: white;
    font-size: 1rem;
    text-align: center;
    padding: 0px;
    margin: 5px;
    cursor: pointer;
    border-radius: 5px;
}

button:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

.logo {
    margin-bottom: 10px;
    cursor: pointer;
}

</style>

