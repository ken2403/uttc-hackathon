import { FC, useState, ChangeEvent, KeyboardEvent } from "react";
import { useHistory } from "react-router-dom";
import axios, { AxiosResponse, AxiosError } from "axios";
import { BaseContainer } from "../atoms/container/BaseContainer";
import { NoUserAlert } from "./NoUserAlert";
import { InputButton } from "../molecules/InputButton";
import { URLs } from "../../api/urls";
import { ResponseGetUser } from "../../types/response";

export const Login: FC = () => {
  const [inputText, setInputText] = useState<string>("");
  const [userName, setUserName] = useState<string>("");
  const [userExist, setUserExist] = useState<boolean>(true);

  const history = useHistory();

  const onChageInput = (event: ChangeEvent<HTMLInputElement>) => {
    setInputText(event.target.value);
  };
  const onClickEnter = () => {
    if (inputText === "") {
      return;
    }
    let url = URLs.getUser;
    url += inputText;
    axios
      .get(url)
      .then((res: AxiosResponse) => {
        if (res.data.length === 0) {
          setUserName(inputText);
          setUserExist(false);
          setInputText("");
        } else {
          const responseUser: ResponseGetUser[] = res.data;
          const newUserName = responseUser[0].user_name;
          setUserName(responseUser[0].user_name);
          setUserExist(true);
          history.push(`/channelindex/${newUserName}`, {
            userName: newUserName
          });
        }
      })
      .catch((error: AxiosError<{ error: string }>) => {
        console.log(error);
      });
  };
  const onEnterPress = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      onClickEnter();
    }
  };

  // rendering
  return (
    <>
      {userExist ? (
        <BaseContainer>
          <h2>ユーザー名を入力してください</h2>
        </BaseContainer>
      ) : (
        <NoUserAlert userName={userName} />
      )}
      <BaseContainer>
        <InputButton
          value={inputText}
          placeholder="put your name"
          onChange={onChageInput}
          onClick={onClickEnter}
          onKeyPress={onEnterPress}
        >
          Enter
        </InputButton>
      </BaseContainer>
    </>
  );
};
