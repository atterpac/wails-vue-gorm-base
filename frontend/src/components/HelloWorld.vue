<script lang="ts" setup>

const otherColors = [
    "primary-color",
    "primary-dark",
    "secondary-color",
    "secondary-dark",
    "ternary-color",
    "ternary-dark",
    "quaternary-color",
    "quaternary-dark",
    ]

const colors = [
    "red",
    "green",
    "yellow",
    "orange",
    "blue",
    "purple",
    "pink",
]

const darkColors = [
    "red-dark",
    "green-dark",
    "yellow-dark",
    "orange-dark",
    "blue-dark",
    "purple-dark",
    "pink-dark",
]
</script>

<template>
    <section>
        <div class="color-swatches">
            <div>
                <div class="colors" v-for="color in otherColors" :style="`background-color: var(--${color})`"> {{ color }} </div>
            </div>
            <div>
                <div class="colors" v-for="color in colors" :style="`background-color: var(--${color})`"> {{ color }} </div>
            </div>
            <div>
                <div class="colors" v-for="color in darkColors" :style="`background-color: var(--${color})`"> {{ color }} </div>
            </div>
        </div>
    </section>
</template>

<style scoped>
section {
    background-color: var(--background-color-light);
    display: flex;
    align-items: center;
    justify-content: center;
}

.color-swatches {
    display: flex;
    align-items: center;
    justify-content: space-between;
    overflow: hidden;
}

.colors {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 50px;
    width: 200px;
    border-radius: 5px;
    font-size: 1.5rem;
    font-weight: bold;
    margin: 20px;
}

.color-wrapper {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}

.color-wrapper > div {
    margin: 10px;
    width: 100px;
    height: 100px;
    border-radius: 5px;
}
</style>
