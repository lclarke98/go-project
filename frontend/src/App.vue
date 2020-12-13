<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Login</h1>

        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
            <input v-model="username" type="text" id="website-input" placeholder="Enter a username" class="form-control">
            <input v-model="password" type="text" id="website-input" placeholder="Enter a password" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'App',

  data() { return {
    username: '',
    password: '',
    thumbnailUrl: '<img :src="thumbnailUrl"/>',
  } },

  methods: {
    makeWebsiteThumbnail() {
  // Call the Go API, in this case we only need the URL parameter.
  axios.post("http://localhost:3000/api/register", {
    username: this.username,
    password: this.password,
  })
  .then((response) => {
    this.thumbnailUrl = response.data.screenshot;
  })
  .catch((error) => {
    window.alert(`The API returned an error: ${error}`);
  })
}
  }
}
</script>