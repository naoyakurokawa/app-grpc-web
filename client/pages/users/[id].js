import { useEffect,useState } from 'react'
import { useRouter } from 'next/router'
import {GetUserByIdRequest, DeleteUserRequest} from '../../lib/hello_pb';
import {HelloServiceClient} from '../../lib/HelloServiceClientPb';
import { UserDetail } from "../../components/UserDetail"
import Link from 'next/link'
import Router from 'next/router'

export default function User() {
  const router = useRouter();
  const [id, setId] = useState(0);
  const [userData, setUserData] = useState('');
  //ユーザー詳細取得
  const getGetUserData = async (id) => {
    const request = new GetUserByIdRequest();
    request.setId(id);
    const client = new HelloServiceClient("http://localhost:8080");
    const response = await client.getUserById(request, {});
    const resUserData = response.toObject().userList;
    setUserData(resUserData);
  }
  //ユーザー削除
  const deleteUser = async (e) => {
    var result = window.confirm('削除してもよろしいですか？');
    if(result){
      const request = new DeleteUserRequest();
      request.setId(id);
      const client = new HelloServiceClient("http://localhost:8080");
      const response = await client.deleteUser(request, {});
      console.log(response);
      Router.push('/')
      return
    }
  }
  useEffect(()=>{
    // idがqueryで利用可能になったら処理される
    if (router.asPath !== router.route) {
      setId(router.query.id);
    }
  },[router])
  useEffect(()=>{
    if (id) {
      getGetUserData(id);
    }
  },[id])
  return (
    <div>
      <UserDetail UserData={userData}/>
      <Link href="/">
        <a>トップページへ戻る</a>
      </Link>
      <br/>
      <button onClick={deleteUser}>削除する</button>
    </div>
  )
}