<template>
    <section>
        <div class='container'>
            <div class='logo'>
                <img src='../../assets/images/logo-universal.png'/>
            </div>
            <div class='login'>
                <h1>Register</h1>
                <input type='text' placeholder='New Password'/>
                <input type='password' placeholder='Confirm Password'/>
                <button class='login-btn'>Reset Password</button>  
                <button class='btn' @click="onClick('/login')">Login</button>
                <button class='btn' @click="onClick('/register')">Register</button>
            </div>
        </div>
    </section>
</template>

<script lang="ts" setup>
import { router } from '../../router';

const onClick = (path: string) => {
    console.log('clicked');
    router.push(path);
}
</script>

<style scoped>
section {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: row;
}

.container {
    display: flex;
    width: 75%;
    height: 75%;
    border: 0.5px solid var(--gray-1);
    border-radius: 10px;
    box-shadow: 0px 0px 20px 0px hsla(var(--hue), 70%, 20%, 0.75);
}

.logo {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background-color: var(--background-color);
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
    width: 30%;
}

img {
    width: 95%;
}

.login {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;
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
    height: 36px;
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
    height: 48px;
    margin: 5px;
    border-radius: 5px;
    border: none;
    outline: none;
    padding: 5px;
    cursor: pointer;
}

</style>

