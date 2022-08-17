import { FC, useEffect, useState, KeyboardEvent, ChangeEvent } from "react";
import { useLocation, useHistory } from "react-router-dom";
import axios, { AxiosResponse, AxiosError } from "axios";

import { URLs } from "../../api/urls";
import { Channel } from "../../types/channel";
import { Message } from "../../types/message";
import { PostMessage } from "../../types/post";
import { User } from "../../types/user";
import { BaseContainer } from "../atoms/container/BaseContainer";
import { SecondaryButton } from "../atoms/button/SecondaryButton";
import { ShowMessage } from "../template/ShowMessage";
import { Page404 } from "./Page404";
import { ResponseGetMessage } from "../../types/response";
import { SendMessagArea } from "../organism/SendMessageArea";

type Props = {
  channel: Channel | null;
};

type State =
  | undefined
  | {
      channel: Channel;
      user: User;
    };

export const Chat: FC<Props> = () => {
  const { state } = useLocation<State>();
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputText, setInputText] = useState<string>("");
  const [getMessage, setGetMessage] = useState<boolean>(true);
  const history = useHistory();

  const onChageInput = (event: ChangeEvent<HTMLInputElement>) => {
    setInputText(event.target.value);
  };

  const onClick = () => {
    if (state===undefined){
      return 
    }else{
    const url = URLs.postMessage;
    var now = new Date();
    var Year = now.getFullYear();
    var Month = now.getMonth() + 1;
    var Dates = now.getDate();
    var Hour = now.getHours();
    var Min = now.getMinutes();
    var Sec = now.getSeconds();
    const date = `${Year}-${Month}-${Dates}-${Hour}:${Min}:${Sec}`;

    const message: PostMessage = {
      message_send_user_id: state?.user.userId,
      message_send_date: date,
      message_channel_id: state?.channel.channelId,
      message_edit_flag: false,
      message_message: inputText
    };
    const json = JSON.stringify(message);

    const headers = {
      "Content-Type": "application/json"
    };

    axios.post(url, json, { headers: headers }).then((res) => {
      console.log(res);
      setGetMessage(!getMessage);
      setInputText("");
    });
    // .catch((err) => {
    //   console.log(err);
    // });
  }
  };

  const onKeyPress = (event: KeyboardEvent<HTMLInputElement>) => {
    onClick();
  };

  useEffect(() => {
    if (typeof state !== "undefined") {
      let url = URLs.getMessage;
      url += state?.channel.channelName;
      const channelMessages: Message[] = [];
      axios
        .get(url)
        .then((res: AxiosResponse) => {
          if (res.data.length !== 0) {
            const responseMessage: ResponseGetMessage[] = res.data;
            responseMessage.map((message: ResponseGetMessage) => {
              channelMessages.push({
                messageId: message.message_id,
                messageSendUserId: message.message_send_user_id,
                messageSendTime: message.message_send_time,
                messageChannelId: message.message_channel_id,
                messageEditFlag: message.message_edit_flag,
                messageMessage: message.message_message
              });
            });
            setMessages(channelMessages);
            console.log(responseMessage);
          }
        })
        .catch((error: AxiosError<{ error: string }>) => {
          console.log(error);
        });
    }
  }, [state, getMessage]);

  const onClickBack = () => {
    history.goBack();
  };

  // rendering
  if (typeof state === "undefined") {
    return <Page404 />;
  } else {
    return (
      <>
        <BaseContainer>
          <h1>{state?.channel.channelName}</h1>
          <ShowMessage messages={messages} userId={state?.user.userId} />
          <SendMessagArea
            onClick={onClick}
            onKeyPress={onKeyPress}
            onChange={onChageInput}
            value={inputText}
          />
        </BaseContainer>
        <SecondaryButton onClick={onClickBack}>
          チャンネル一覧に戻る
        </SecondaryButton>
      </>
    );
  }
};
