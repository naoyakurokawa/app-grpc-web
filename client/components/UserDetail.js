export const UserDetail = ({UserData}) => {
  return(
    <div>
      {
        Object.keys(UserData).map((key) => {
          return (
            <div  key={UserData[key].id}>
              名前：{UserData[key].name}<br/>
              スコア：{UserData[key].score}<br/>
              画像：<br/><img src = {UserData[key].photourl}/>
            </div>
          )
        })
      }
    </div>
  )
}