<script>
  import Home from "./Home.svelte";
  import Login from "./Login.svelte";
  import { Button } from "sveltestrap";
  import { Alert } from 'sveltestrap';
  import axios from 'axios';
  import Cookies from 'js-cookie';
  
  
  const color = "info";
  const warning = "danger";
  const btncolor = "light";
  let visible = true;
  let page;
  let User = { isreg : false , username : "" , isexist : false, already:false};
  
  // CheckCookie checks for cookie if it available or not
  function CheckCookie() {
    var c = Cookies.get("session");
    if (c != null) {
      User.already = true;
    } else {
      User.already = false;
    }
  }

CheckCookie()

  //check function checks if the the passwords fields are identical or not
  function check() {
    var password = document.getElementById("pw1");
    var repassw = document.getElementById("pw2");
    var btn = document.getElementById("sub");
    if (password.value !== repassw.value){
      btn.disabled = true;
    } else {
      btn.disabled = false;
    }
  }

  function showP1(){
    var pw1 = document.getElementById("pw1");
    if (pw1.type === "password"){
      pw1.type = "text";
    } else {
      pw1.type = "password";
    }
  }

  function showP2(){
    var pw2 = document.getElementById("pw2");
    if (pw2.type === "password") {
      pw2.type = "text";
    } else {
      pw2.type = "password";
    }
  }

  //register function for makes a post request to our api then check if the user is registred or not
  function register(){ 
    var user = document.getElementById("user");
    var password = document.getElementById("pw1");
    var repassw = document.getElementById("pw2");
    var email = document.getElementById("email");
    //`user=${user.value}&email=${email.value}&passw=${password.value}`,
    if (/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email.value)){
      User.username = user.value
      var bodyFormData = new FormData();
      bodyFormData.set("user",user.value);
      bodyFormData.set("passw",password.value);
      bodyFormData.set("email",email.value);

      axios({
        method:'post',
        url:'http://localhost:8080/register',
        data:bodyFormData,
        reponseType: 'json',
        headers:{ 'Content-Type':'multipart/form-data' }
        })
        .then(function (response) {
          console.log(response.status)
          var resp = response.data;
          var r = JSON.stringify(resp);
          var obj = JSON.parse(r);
          User.isreg = obj.IsReg;
          if(User.isreg){
            window.location.replace("/login")
          } else {
            User.isexist = true
          }
          })
        .catch(function (error) {
          console.log(JSON.stringify(error.response))
        });
      } 
    user.value     = "";
    password.value = "";
    repassw.value  = "";
    email.value    = "";
  }

  function subm(e) {
    e.preventDefault();
    register();
    if (!visible) {
      visible = true
    }
  }

</script>

<main>
  { #if User.isexist }
    <Alert color={color} isOpen={visible} toggle={() => (visible = false)}>
       {User.username}, Already Exist 
    </Alert>
  { /if }
  { #if User.already }
    <Alert color={warning} isOpen={visible} toggle={() => (visible = false)}>
      You are Already logged in
    </Alert>
  { :else }
  <div class="register_field">
    <form method="post" id="myForm" on:submit={subm}>
      <input id="user" required={true} type="text" placeholder="Username" name="user">
      <input id="pw1" on:keyup={check} required={true} type="password" placeholder="password" name="passw">
      <span on:click={showP1}>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path  d="M12.015 7c4.751 0 8.063 3.012 9.504 4.636-1.401 1.837-4.713 5.364-9.504 5.364-4.42 0-7.93-3.536-9.478-5.407 1.493-1.647 4.817-4.593 9.478-4.593zm0-2c-7.569 0-12.015 6.551-12.015 6.551s4.835 7.449 12.015 7.449c7.733 0 11.985-7.449 11.985-7.449s-4.291-6.551-11.985-6.551zm-.015 3c-2.209 0-4 1.792-4 4 0 2.209 1.791 4 4 4s4-1.791 4-4c0-2.208-1.791-4-4-4z"/></svg>
      </span>
      <input id="pw2" on:keyup={check} required={true} type="password" placeholder="Retype your password" name="passw2">
      <span on:click={showP2}>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path  d="M12.015 7c4.751 0 8.063 3.012 9.504 4.636-1.401 1.837-4.713 5.364-9.504 5.364-4.42 0-7.93-3.536-9.478-5.407 1.493-1.647 4.817-4.593 9.478-4.593zm0-2c-7.569 0-12.015 6.551-12.015 6.551s4.835 7.449 12.015 7.449c7.733 0 11.985-7.449 11.985-7.449s-4.291-6.551-11.985-6.551zm-.015 3c-2.209 0-4 1.792-4 4 0 2.209 1.791 4 4 4s4-1.791 4-4c0-2.208-1.791-4-4-4z"/></svg>
      </span>
      <input id="email" required={true} type="email" placeholder="email" name="email"><br>
      <Button name="btn" id="sub" {btncolor}>Register</Button>
    </form>	
  </div>
  { /if }
</main>
<style>
  span{
    position: absolute;
    right: 26%;
    padding-top : 1.5%;
  }

	.register_field{
		padding: 16px;
    padding-top:5%;
	}
	input[type=text], input[type=password],input[type=email] {
  width: 50%;
  padding: 15px;
  margin: 5px 0 22px 0;
  display: inline-block;
  border: none;
  background: #f1f1f1;
	}
	input[type=text]:focus, input[type=password]:focus, input[type=email]:focus {
  background-color: #ddd;
  outline: none;
	}
	.sub {
		background-color: #8cd7f4;
		color: black;
		padding: 16px 20px;
		margin: 8px 0;
		border: none;
		cursor: pointer;
		width: 50%;
		opacity: 0.9;
	}
	.sub:hover {
  	opacity:1;
	}
</style>

