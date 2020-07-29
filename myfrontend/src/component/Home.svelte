<script>
import axios from 'axios';
import Cookies from 'js-cookie';
//import { setContext } from 'svelte';
import { Alert, Button, Card, CardBody, CardFooter, CardHeader,
  CardImg, CardSubtitle, CardText, CardTitle} from "sveltestrap";

const cookiename = "session";
let visible = true;
let color = "success" ;
let User;

// GetUserByCookie Check if there is a cookie if there is one 
// we send a GET request with our cookie in the header 
// to get the cookie decrypted to get the username of the user
// if there is not cookie we show Guest as a username of the user.
function GetUserByCookie(){
  var c = Cookies.get(cookiename);
  if (c != null) {
    axios({
      method:'get',
      url:'http://localhost:8080/user',
      responseType:'json',
      withCredentials:true
      })
      .then(function (response) {
        var resp = response.data;
        var r = JSON.stringify(resp);
        var obj = JSON.parse(r);
        var userName = obj.Username;
        User = userName;
      })
  } else {
    User = "Guest";
  }
} 

GetUserByCookie()


</script>

<main>
  { #if User != "Guest" }
    <Alert color={color} isOpen={visible} toggle={() => (visible = false)} >
      Welcome, {User} Enjoy your Visit!
    </Alert>
  { /if }
  <h2>Welcome {User}</h2>

  <div class="cred-Card">
    <Card>
      <CardHeader>
        <CardTitle>Credentials</CardTitle>
      </CardHeader>
      <CardBody>
        <!--><CardSubtitle>Card subtitle</CardSubtitle><-->
        <CardText>
          <a href="/credentials"> Access You Creds </a>
        </CardText>
      </CardBody>
    </Card>
  </div>
  <div class="upload-card">
    <Card>
      <CardHeader>
        <CardTitle>Upload a Csv file</CardTitle>
      </CardHeader>
      <CardBody>
        <CardSubtitle>Upload a Csv file contains your credentials.</CardSubtitle>
        <CardText>
          <a href="/credfile"> Upload CSV FILE </a>
        </CardText>
      </CardBody>
    </Card>
  </div>
  <div class="gen-pw">
    <Card>
      <CardHeader>
        <CardTitle>Generate A Password</CardTitle>
      </CardHeader>
      <CardBody>
        <CardText>
          <a href="/generatepw"> Generate Password </a>
        </CardText>
      </CardBody>
    </Card>
  </div>

</main>

<style>

  .cred-Card {
    margin-left:4%;
    margin-top:6%;
    width:20%;
    float:left;
  }

  .upload-card {
    margin-top:6%;
    margin-left:8%;
    width:35%;
    float:left;
  }

  .gen-pw {
    margin-left:11%;
    width:20%;
    margin-top:6%;
    float:left;
  }

  h2{
    padding-top:5%;
    margin-bottom:5%;
  }

</style>
