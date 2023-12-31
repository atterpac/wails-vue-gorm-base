<template>
    <section>
        <div class='logo'>
            <img src='../../assets/images/logo-universal.png'/>
        </div>
        <div class='login'>
            <h1>Login</h1>
            <input type='text' placeholder='Username'/>
            <input type='password' placeholder='Password'/>
            <button class='login-btn'>Login</button>
            <button class='btn' @click="onClick('/register')">Register</button>
            <button class='btn'>Reset Password</button>
        </div>
    </section>
</template>

<script lang="ts" setup>
import { router } from '../../router';
// import { User } from '../../../wailsjs/go/models';

const onClick = (path: string) => {
    console.log('clicked');
    router.push(path);
}
</script>

<style scoped>
section {
    background-color: var(--background-color-light);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: row;
    height: 95%;
    width: 95%;
    box-shadow: 0px 0px 20px 0px hsla(var(--hue), 70%, 20%, 0.75);
}

.logo {
    display: flex;
    flex-direction: column;
    flex: 1;
    align-items: center;
    justify-content: center;
    background-color: var(--background-color);
    height: 100%;
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
}

img {
    height: 30%;
}

.login {
    display: flex;
    flex-direction: column;
    flex: 2;
    align-items: center;
    justify-content: center;
    width: 50%;
    height: 50%;
    background-color: var(--background-color-light);
    border-radius: 5px;
}

.btn {
    background-color: var(--gray-2);
    color: var(--text-color);
}

.btn:hover {
    background-color: var(--gray-3);
}

.login-btn {
    background-color: var(--primary-dark);
    color: var(--text-color-light);
}

.login-btn:hover {
    background-color: var(--primary-color);
    color: var(--text-color);
}

input {
    width: 50%;
    height: 10%;
    margin: 10px;
    border-radius: 5px;
    background-color: var(--background-color);
    color: var(--text-color);
    border: none;
    outline: none;
    padding: 5px;
    padding-left: 10px;
}

button {
    width: 30%;
    height: 10%;
    margin: 5px;
    border-radius: 5px;
    border: none;
    outline: none;
    padding: 5px;
    cursor: pointer;
}
</style>

