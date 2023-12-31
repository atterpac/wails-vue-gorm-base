<template>
    <section class="nav-bar">
        <h1>Wails Example</h1>    
        <button 
            v-for="item in navigationItems" 
            :key=item.path
            @click='onClick(item.path)'>
            {{item.name}}
        </button>
    </section>
</template>

<script lang="ts" setup>
import { router } from '../../router';

type NavigationItem = {
    name: string;
    path: string;
}

const navigationItems: NavigationItem[] = [
    {
        name: 'Home',
        path: '/'
    },
    {
        name: 'Auth',
        path: '/login'
    },
    {
        name: 'Runtime',
        path: '/runtime'
    },
]

const onClick = (path: string) => {
    router.push(path);
}
</script>

<style scoped>
.nav-bar {
    position: fixed;
    left: 0;
    top: 36px;
    margin-top: 24px;
    height: 100%;
    width: 10%;
    z-index: 100;
}

button {
    width: 85%;
    height: 36px;
    border: none;
    background-color: transparent;
    color: white;
    font-size: 1rem;
    text-align: center;
    padding: 0px;
    margin: 10px;
    cursor: pointer;
    border-radius: 5px;
}

button:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

</style>

