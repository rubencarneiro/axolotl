<template>
  <div id="app">
    <header-comp></header-comp>
    <div class="container">
      <error-modal  v-if="error"/>
      <router-view />
    </div>
  </div>
</template>

<script>
window.getCookie = function(cname) {
      var name = cname + "=";
      var ca = document.cookie.split(';');
      for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
          c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
          return c.substring(name.length, c.length);
        }
      }
      return false;
}
if (window.getCookie("darkMode") === 'true') {
  import ('./assets/dark.scss');
} else {
  import ('./assets/light.scss');
}
import HeaderComp from "@/components/Header.vue"
import ErrorModal from "@/components/ErrorModal.vue"
export default {
  name: 'axolotl-web',
  components: {
    HeaderComp,
    ErrorModal
  },
  mounted(){
    var userLang = navigator.language || navigator.userLanguage;
    this.$language.current = userLang;
  },
  computed: {
    error () {
      return this.$store.state.error
    }
  }
}
</script>
<style>
::-webkit-scrollbar {
    display: none;
}
#app{
  font-family:"ubuntu"
}
#app >.container{
  padding-top: 50px;
  position:relative;
}
.btn:focus{
  box-shadow:none;
}
.btn{
  border-radius:0px;
}
.btn-primary {
  background-color: #2090ea;
}
.no-entries {
    height: 90vh;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
