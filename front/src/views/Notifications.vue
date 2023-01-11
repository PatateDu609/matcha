<template>
    <br>
    <div class="notif-container">
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
        <div class="notif">
            <p>Ceci est une notification !</p>
        </div>
    </div>
</template>

<style>
.notif-container {
    top: 10%;
    width: auto;
    height: 75vh;
    padding: 20px;
    position: relative;
    overflow: scroll;
}

.notif {
    display: flex;
    flex-direction: column;
    background-color: white;
    border-radius: 10px;
    padding: 20px;
    color: black;
    font-size: 20px;
    font-weight: bold;
    margin: 10px;
}
</style>