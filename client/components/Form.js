import { useState } from 'react';
import {HelloRequest} from '../lib/hello_pb';
import {HelloServiceClient} from '../lib/HelloServiceClientPb';
export const Form = ( {onAddData} ) => {
  const [text, setText] = useState('');
  const submitForm = async(event) => {
    event.preventDefault();
    //grpcにformの入力データを渡す
    const request = new HelloRequest();
    request.setName(text);
    const client = new HelloServiceClient("http://localhost:8080");
    const response = await client.sayHello(request, {});
    onAddData(response.getMessage());
  }
  return (
  <div>
    <h4>grpc-test-form</h4>
    <form onSubmit={submitForm}>
      <div>
        <input
          type ="text"
          value={text}
          onChange={(e)=>setText(e.target.value)}
        />
      </div>
      <div>
        <button>送信</button>
      </div>
    </form>
  </div>
  )
}