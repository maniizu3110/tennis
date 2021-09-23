import React from "react";
import axios from 'axios'

const App = () => {
    const CreateUser = async () => {
        //TODO:非同期処理はHooksで実装
        await axios.post("http://localhost/api/v1/sign/").then(res=>{
          console.log(res);
        })
      };
  

  return (
    <div>
      <button onClick={CreateUser}>ユーザー作成</button>
    </div>
  );
};

export default App;
