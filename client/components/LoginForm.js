import { useState } from 'react';
import {LoginRequest} from '../lib/hello_pb';
import {HelloServiceClient} from '../lib/HelloServiceClientPb';
import Router from 'next/router';
import { useCookies } from 'react-cookie';

export const LoginForm = () => {
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [cookies, setCookie] = useCookies(['login_token']);

  //フォームsubmit時の処理
  const submitLoginForm = async(event) => {
    event.preventDefault();
    const request = new LoginRequest();
    request.setName(name);
    request.setPassword(password);
    const client = new HelloServiceClient("http://localhost:8080");
    const response = await client.login(request, {});
    console.log("レスポンス",response);
    if(response.toObject().islogin){
      setCookie('login_token', response.toObject().token, { path: '/' });
      console.log("aaaaaa",cookies);
      Router.push('/users/list')
      return
    }
  }
  return (
    <div className="user-form-container">
      <h4>ログイン</h4>
      <form onSubmit={submitLoginForm}>
        <div>
          <p>ユーザー名</p>
          <input
            type ="text"
            onChange={(e)=>setName(e.target.value)}
          />
        </div>
        <div>
          <p>パスワード</p>
          <input
            type ="password"
            onChange={(e)=>setPassword(e.target.value)}
          />
        </div>
        <div>
          <button>送信</button>
        </div>
      </form>
      <style jsx>{`
      .user-form-container {
        margin:10px;
      }
      .profile-image-sec{
        margin:10px;
      }
      .profile-image {
        width: 100px;
        height: 100px;
      }
      `}
      </style>
    </div>
  )
}