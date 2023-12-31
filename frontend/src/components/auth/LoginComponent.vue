<template>
    <section>
        <div class='logo'>
            <img src='../../assets/images/logo-universal.png'/>
        </div>
        <div class='login'>
            <h1>Login</h1>
            <input type='text' placeholder='Username'/>
            <input type='password' placeholder='Password'/>
            <button>Login</button>
            <div>or</div>
            <button @click="onClick('/register')">Register</button>
            <button>Reset Password</button>
        </div>
    </section>
</template>

<script lang="ts" setup>
import { router } from '../../router';
import { User } from '../../../wailsjs/go/models';

const onClick = (path: string) => {
    console.log('clicked');
    router.push(path);
}
</script>

<style scoped>
section {
    background-color: #282c34;
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: row;
    height: 95%;
    width: 95%;
}

.logo {
    display: flex;
    flex-direction: column;
    flex: 1;
    align-items: center;
    justify-content: center;
    background-color: #383c44;
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
    background-color: #282c34;
    border-radius: 5px;
}

input {
    width: 50%;
    height: 10%;
    margin: 10px;
    border-radius: 5px;
    border: none;
    outline: none;
    padding: 5px;
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

