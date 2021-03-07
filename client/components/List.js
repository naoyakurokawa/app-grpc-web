import { useEffect} from 'react';

//親コンポーネント(App.js)から受け取ったプロパティ(langs)を使用
export const List = ({formData}) => {
  useEffect(()=>{
    //mouting,updateで呼ばれる
    // console.log('List.js:useEffect');
    //unmountで呼ばれる
    return()=>{
      // console.log('List.js:useEffect:ummount');
    }
  })
  return(
    <div>
      {
        formData.map((data,index) => {
          return <div key={index}>{data}</div>
        })
      }
    </div>
  )
}