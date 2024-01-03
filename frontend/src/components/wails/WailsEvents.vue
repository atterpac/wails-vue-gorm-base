<template>
    <section>
        <div class='header'>
            <h1>Events On</h1>
            <button class='timer'>Time: {{time}}</button>
        </div>
        <p>This method sets up a listener for the given event name. When an event of type eventName is emitted, the callback is triggered. Any additional data sent with the emitted event will be passed to the callback. It returns a function to cancel the listener.</p>
            <div class='examples'>
                <div class='example'>
                    <CodeBlock class='code-block' code_height='250px' :code="timeSetup" lang='go' label='Golang'/>
                </div>
                <div class='example'>
                    <CodeBlock class='code-block' code_height='250px' :code="timeRecieve" lang='TypeScript' label='TypeScript'/>
                </div>
            </div>
    </section>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import CodeBlock from './CodeBlock.vue'
import { EventsOn } from '../../../wailsjs/runtime/runtime';

const time = ref(0);
EventsOn('app:tick', (data: any) => {
    time.value = data;
});

const timeSetup = `
// import "github.com/wailsapp/wails/v2/pkg/runtime"
func (a *App) startTimerEmit() {
	for {
		timeNow := time.Now().Format("15:04:05")
		runtime.EventsEmit(a.ctx, "app:tick", timeNow)
		time.Sleep(1 * time.Second)	
	}
}`

const timeRecieve = `
import { EventsOn } from '../wailsjs/runtime/runtime';

const time = 0;

EventsOn('app:tick', (data: any) => {
    time.value = data;
});`
</script>

<style scoped>

section {
    background-color: var(--background-color-light);
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.header {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 48px;
    align-items: center;
    justify-content: space-around;
}

.timer {
    width: 20%;
    height: 32px;
    background-color: var(--primary-color);
    color: var(--text-color);
    border: 0.5px solid var(--gray-1);
    border-radius: 10px;
}

h1 {
    color: var(--text-color);
    margin-bottom: 20px;
    font-size: 1.5rem;
    justify-self: flex-start;
    padding-left: 20px;
    text-align: left;
    margin-bottom: 5px;
}

.examples {
    display: flex;
    width: 100%;
    height: 250px;
    align-items: center;
    justify-content: space-around;
    flex-direction: row;
}

.example {
    display: flex;
    width: 45%;
    height: 200px;
    align-items: center;
    justify-content: center;
    flex-direction: column;
}

</style>
