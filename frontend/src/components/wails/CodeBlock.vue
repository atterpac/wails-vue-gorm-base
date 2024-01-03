<template>
    <section class='codeblock'>
        <div class="code">
            <pre>
                <code :class="`language-`+props.lang" v-html='highlight' />
            </pre>
            <div class='lang'>{{props.label}}</div>
        </div>
    </section>
</template>

<script lang="ts" setup>
import hljs from 'highlight.js/lib/core';
import 'highlight.js/styles/atom-one-dark.css';
import go from 'highlight.js/lib/languages/go';
import typescript from 'highlight.js/lib/languages/typescript';

const props = defineProps({
    code: {
        type: String,
        required: true
    },
    label: {
        type: String,
        default: 'Code'
    },
    lang: {
        type: String,
        default: 'go'
    },
    code_width: {
        type: String,
        default: '100%'
    },
    code_height: {
        type: String,
        default: '100%'
    }
})

hljs.registerLanguage('go', go);
const highlight = hljs.highlight('go', props.code).value;

hljs.highlightAll();
</script>

<style>
.codeblock {
    background-color: var(--background-color-light);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    width: 50%;
    text-align: start;
}

pre {
    tab-size: 2;
    padding-left: 20px;
}

.code {
  width: 100%;
  height: 200px;
  background-color: var(--background-color-light);
  border: 1px solid var(--gray-4);
  border-radius: 4px;
  padding: 5px;
  overflow-x: auto;
  position: relative;
}

.lang {
    position: absolute;
    top: 0; 
    right: 0; 
    padding-top: 0.25rem;
    padding-bottom: 0.25rem; 
    padding-left: 0.5rem;
    padding-right: 0.5rem;
    border-bottom-left-radius: 0.375rem; 
    font-size: 0.75rem;
    line-height: 1rem; 
    font-weight: 700; 
    text-transform: uppercase; 
}
</style>

