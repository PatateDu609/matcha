
<template>
<br><br>
<div class="UserList">
  <div class="User">
    <img alt="Prénom Nom" class="ProfilePicture" src="https://pbs.twimg.com/media/DqOLIIgWsAIG6x9?format=jpg&name=4096x4096" />
    <div class="centered">John Doe</div>
    <div class="buttons">
      <button class="check">❌</button>
      <button class="check">✅</button>
    </div>
  </div>
</div>
</template>

<style>

.check {
  color: #fff;
  margin: 10px;
  padding: 5px;
  font-size: 31px;
  height: 75px;
  width: 75px;
  border-radius: 70%;
  border: none;
  box-shadow: 0 2px 4px darkslategray;
  cursor: pointer;
  transition: all 0.2s ease;
}

.check:active {
  box-shadow: 0 0 2px darkslategray;
  transform: translateY(2px);
}

.UserList {
  all: unset;
  top: 20vh;
  left: 10%;
  position: absolute;
  width: 100%;
  height: 100%;
}

.ProfilePicture {
  max-width: 100%;
  max-height: 100%;
  border-radius: 25px;
}

.User {
  width: 1000px;
  height: 500px;
  padding: 20px;
  position: relative;
  text-align: center;
  color: white;
}

.User > .centered {
  position: absolute;
  top: 75%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  color: black;
  border-radius: 10px;
  padding: 10px;
  font-size: 20px;
  font-weight: bold;
}

.User > .buttons {
  display: flex;
  flex-direction: row;
  align-items: center;
  left: 40%;
}
</style>