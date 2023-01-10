import { RouterLink } from 'vue-router';

<template>
    <header>
        <nav id="main-menu">
            <ul class="nav-bar">
                <img alt="bubbler logo" src="@/assets/logo.svg" width="50" height="50" />
                <li><RouterLink to ="/">Home</RouterLink></li>
                <li><RouterLink to ="/matchs">Matchs</RouterLink></li>
                <li><RouterLink to ="/search">Search</RouterLink></li>
                <li><RouterLink to ="/notifications">Notifications</RouterLink></li>
                <li><RouterLink to ="/profile">Profile</RouterLink></li>
                <li><RouterLink to ="/login">Login</RouterLink></li>
                <li><RouterLink to ="/register">Register</RouterLink></li>
            </ul>
        </nav>
    </header>
</template>

<style scoped>
#main-menu .nav-bar { top: 20px; }
#main-menu .nav-bar li { display: inline; padding:0 10px; }
#main-menu img { position: absolute; left: 0; top: -20px;}
#main-menu .nav-bar li a { padding-left: 25px; text-transform: uppercase; }
</style>