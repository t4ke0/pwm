<script>

import { Alert } from "sveltestrap";
import { Button } from "sveltestrap";
import axios from 'axios';
import Cookies from 'js-cookie';



//import https from 'https';
//import fs from 'fs';
//import { USERNAME } from './store';

//const certPath = "/home/takeo/Desktop/TAKEO/thirdpartie/ssl-proxy/cert.pem"; // put the cert in a convenient way
//      httpsAgent: new https.Agent({ ca: certPath }),

let btncolor = "light";
let visible = true;
let succolor = "success";
let failcolor = "danger";
let ExpireTime = 3600; // maybe export this variable 
let Log = { islog : false , cookieExist: false , wrongCred: false};

const user = { name : "" };

function showP() {
  var pw = document.getElementById("pw1")
  if (pw.type === "password"){
    pw.type = "text";
  } else {
    pw.type = "password";
  }
}

function checkForCookie() {
  var cookie = Cookies.get("session");
  if (cookie != null) {
    return true;
  } else {
    return false ;
  }
}


function login() {
  var username = document.getElementById("user");
  var password = document.getElementById("pw1");
  var bodyFormData = new FormData();

  user.name = username.value;

  bodyFormData.set("user",username.value);
  bodyFormData.set("passw",password.value);

  axios({
      method:'post',
      url:'http://localhost:8080/login',
      data:bodyFormData,
      responseType:'json',
      withCredentials:true,
      headers:{ 'Content-Type': 'multipart/form-data' }
    })
    .then((response) => {
        var resp = response.data;
        var r = JSON.stringify(resp);
        var obj = JSON.parse(r);
        var islog = obj.IsLog;

        if (islog) {
          Log.islog = true ;
          if (checkForCookie) {
            window.location.replace("/");//replace "home dir" to credentials path
          }
        } else {
            Log.wrongCred = true ;
        }
      })
  username.value = "";
  password.value = "";
}


function AlreadyLogged() {
  var c = Cookies.get("session");
  if (c != null) {
    Log.cookieExist = true;
  } else {
    Log.cookieExist = false;
  }
}

AlreadyLogged()

function submt(e) {
  e.preventDefault();
  login()
}
</script>


<main>
  { #if Log.wrongCred }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
       Wrong Username/Password
    </Alert>
  { /if }
  { #if Log.cookieExist }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      You are Already Loggedin 
    </Alert>
  { :else }
  <div class="login_field">
    <form method="post" id="myForm" on:submit={submt}>
      <input id="user" required={true} type="text" placeholder="Username" name="user">
      <input id="pw1"  required={true} type="password" placeholder="password" name="passw">
      <span on:click={showP}>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path  d="M12.015 7c4.751 0 8.063 3.012 9.504 4.636-1.401 1.837-4.713 5.364-9.504 5.364-4.42 0-7.93-3.536-9.478-5.407 1.493-1.647 4.817-4.593 9.478-4.593zm0-2c-7.569 0-12.015 6.551-12.015 6.551s4.835 7.449 12.015 7.449c7.733 0 11.985-7.449 11.985-7.449s-4.291-6.551-11.985-6.551zm-.015 3c-2.209 0-4 1.792-4 4 0 2.209 1.791 4 4 4s4-1.791 4-4c0-2.208-1.791-4-4-4z"/></svg><br>
      </span><br>
      <Button name="btn" id="sub" {btncolor}>Login</Button>
    </form>	
  </div>
  <a id="link" href="/register">You Are not Registred Yet</a><br>
  <a id="forget" href="/forgotpw">Forgot Your Password ? </a>
  { /if }
</main>

<style>
  #link #forget{
    padding-top :50%;
    text-align: center ;
  }

  span{
    position: absolute;
    right: 26%;
    padding-top : 1.5%;
  }

  .login_field {
    position: relative;
    /*display: inline-block;*/
    padding:16px;
    padding-top:5%;
  }

  input[type=password]::-ms-reveal {
    display: none;
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
      <!--<span>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12.015 7c4.751 0 8.063 3.012 9.504 4.636-1.401 1.837-4.713 5.364-9.504 5.364-4.42 0-7.93-3.536-9.478-5.407 1.493-1.647 4.817-4.593 9.478-4.593zm0-2c-7.569 0-12.015 6.551-12.015 6.551s4.835 7.449 12.015 7.449c7.733 0 11.985-7.449 11.985-7.449s-4.291-6.551-11.985-6.551zm-.015 3c-2.209 0-4 1.792-4 4 0 2.209 1.791 4 4 4s4-1.791 4-4c0-2.208-1.791-4-4-4z"/></svg>
      </span>-->
