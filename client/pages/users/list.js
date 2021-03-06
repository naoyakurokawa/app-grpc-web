import Head from 'next/head'
import { useEffect,useState } from 'react'
import { useRouter } from 'next/router'
import {HelloServiceClient} from '../../lib/HelloServiceClientPb';
import { UserList } from "../../components/UserList"
import Link from 'next/link'
import Router from 'next/router'
import {GetUsersRequest} from '../../lib/hello_pb';
import { useCookies } from 'react-cookie';

export default function UsersList() {
  const [dbData, setDbData] = useState({});
  const [cookies, setCookie, removeCookie] = useCookies(['login_token']);
  const metadata = {'login_token': cookies.login_token}
  //データベースのデータを取得、表示
  const fetchDbData = async ()=>{
      const request = new GetUsersRequest();
      const client = new HelloServiceClient("http://localhost:8080");
      const response = await client.getUsers(request, metadata);
      const userList = response.toObject();
      if(userList.usersList.length == 0){
        Router.push('/')
        return
      }else{
        setDbData(userList["usersList"]);
      }
  }
  const checkIsLogin = ()=>{
    if(!cookies.login_token){
      Router.push('/')
      return
    }
  }
  const logout = () =>{
    removeCookie('login_token', { path: '/' });
    Router.push('/')
  }
  useEffect(()=>{
    checkIsLogin();
    fetchDbData();
  },[])

  return (
    <div className="container">
      <Head>
        <title>Learn Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1 className="title">
          User List
        </h1>

        {/* <Form onAddLang={addLang}/> */}
        <div className="fetch-db-sec">
          <h3>grpc-test-fetch-userTable</h3>
          <UserList userList={dbData}/>
        </div>
        <button onClick={logout}>ログアウト</button>
      </main>

      <footer>
        <a
          href=""
          target="_blank"
          rel="noopener noreferrer"
        >
          © 2021 Learn Next.js
          {/* <img src="/vercel.svg" alt="Vercel Logo" className="logo" /> */}
        </a>
      </footer>

      <style jsx>{`
        .container {
          min-height: 100vh;
          padding: 0 0.5rem;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
        }

        main {
          padding: 5rem 0;
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
        }

        footer {
          width: 100%;
          height: 100px;
          border-top: 1px solid #eaeaea;
          display: flex;
          justify-content: center;
          align-items: center;
        }

        footer img {
          margin-left: 0.5rem;
        }

        footer a {
          display: flex;
          justify-content: center;
          align-items: center;
        }

        a {
          color: inherit;
          text-decoration: none;
        }

        .title a {
          color: #0070f3;
          text-decoration: none;
        }

        .title a:hover,
        .title a:focus,
        .title a:active {
          text-decoration: underline;
        }

        .title {
          margin: 0;
          line-height: 1.15;
          font-size: 4rem;
        }

        .title,
        .description {
          text-align: center;
        }

        .description {
          line-height: 1.5;
          font-size: 1.5rem;
        }

        code {
          background: #fafafa;
          border-radius: 5px;
          padding: 0.75rem;
          font-size: 1.1rem;
          font-family: Menlo, Monaco, Lucida Console, Liberation Mono,
            DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
        }

        .grid {
          display: flex;
          align-items: center;
          justify-content: center;
          flex-wrap: wrap;

          max-width: 800px;
          margin-top: 3rem;
        }

        .card {
          margin: 1rem;
          flex-basis: 45%;
          padding: 1.5rem;
          text-align: left;
          color: inherit;
          text-decoration: none;
          border: 1px solid #eaeaea;
          border-radius: 10px;
          transition: color 0.15s ease, border-color 0.15s ease;
        }

        .card:hover,
        .card:focus,
        .card:active {
          color: #0070f3;
          border-color: #0070f3;
        }

        .card h3 {
          margin: 0 0 1rem 0;
          font-size: 1.5rem;
        }

        .card p {
          margin: 0;
          font-size: 1.25rem;
          line-height: 1.5;
        }

        .logo {
          height: 1em;
        }

        @media (max-width: 600px) {
          .grid {
            width: 100%;
            flex-direction: column;
          }
        }

        .fetch-form-sec{
          border: 2px solid;
          border-color: #031de2;
          padding: 30px;
        }

        .fetch-db-sec{
          border: 2px solid;
          border-color: #031de2;
          padding: 20px;
          margin-top: 10px;
        }

        .register-user-sec{
          border: 2px solid;
          border-color: #031de2;
          padding: 20px;
          margin-top: 10px;
        }
      `}</style>

      <style jsx global>{`
        html,
        body {
          padding: 0;
          margin: 0;
          font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
            Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
            sans-serif;
        }

        * {
          box-sizing: border-box;
        }
      `}</style>
    </div>
  )
}