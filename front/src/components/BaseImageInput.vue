<template>
  <div class="base-image-input" :style="{ 'background-image': `url(${imageData})` }" @click="chooseImage">
    <span v-if="!imageData" class="placeholder"></span>
    <input class="file-input" ref="fileInput" type="file" @input="onSelectFile" >
  </div>
</template>

<script>
export default {
    props: ['number'], // ???
    data () {
        return {
            imageData: null
        }
    },
    methods: {
    chooseImage () {
      this.$refs.fileInput.click()
    },

    onSelectFile () {
      const input = this.$refs.fileInput
      const files = input.files
      if (files && files[0]) {
        const reader = new FileReader
        reader.onload = e => {
          this.imageData = e.target.result
        }
        reader.readAsDataURL(files[0])
        this.$emit('input', files[0])

        let photo = files[0];
        let formData = new FormData();
        formData.append("photo", photo);
        try {
          fetch('http://localhost:4000/upload', {method: "POST", body: formData});
          console.log("image ", number, " uploaded !")
        } catch(e) {
          console.log("upload error: ", e)
        }
      }
    }
  }
}

</script>

<style scoped>
.base-image-input {
  display: inline-block;
  width: 100px;
  height: 100px;
  cursor: pointer;
  background-size: cover;
  background-position: center center;
  padding-left: 10px;
}
.placeholder {
  background: grey;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #333;
  font-size: 18px;
  font-family: Helvetica;
}
.placeholder:hover {
  background: black;
}
.file-input {
  display:none;
}
</style>