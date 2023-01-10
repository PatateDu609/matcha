
<template>
<br><br>
<div class="UserList">
  <div class="User">
    <img alt="Prénom Nom" class="ProfilePicture" src="https://pbs.twimg.com/media/DqOLIIgWsAIG6x9?format=jpg&name=4096x4096" />
  </div>
  <div class="buttons">
    <button class="check">❌</button>
    <button class="check">✅</button>
  </div>
</div>
</template>

<style>
.buttons {
  left: 7%;
  display: flex;
  flex-direction: row;
  align-items: center;
}

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
  left: 40vw;
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
}
</style>