import { FC, useState, ChangeEvent, KeyboardEvent } from "react";
import styled from "styled-components";
import axios from "axios";

import { URLs } from "../../api/urls";
import { BaseContainer } from "../atoms/container/BaseContainer";
import { Modal } from "../atoms/container/ModalContainer";
import { BaseButton } from "../atoms/button/BaseButton";
import { InputButton } from "../molecules/InputButton";
import { SecondaryButton } from "../atoms/button/SecondaryButton";
import { ShowUserName } from "../molecules/ShowUserName";
import { Message } from "../../types/message";
import { PutMessage } from "../../types/put";

type Props = {
  message: Message;
  userId: string;
};

const ChatMessageButton = styled(BaseButton)`
  width: 500px;
  height: 200px;
  border-radius: 10px;
  color: #fff;
  background-color: #aaa;
  font-size: 20px;
`;

export const ShowOneChat: FC<Props> = (props) => {
  const { message, userId } = props;
  const [showModal, setShowModal] = useState<boolean>(false);
  const [inputText, setInputText] = useState<string>(message.messageMessage);

  const onClickMessage = () => {
    if (message.messageSendUserId === userId) {
      setShowModal(true);
    }
  };

  const onChageInput = (event: ChangeEvent<HTMLInputElement>) => {
    setInputText(event.target.value);
  };

  const onClick = () => {
    const url = URLs.postMessage;

    const putMessage: PutMessage = {
      message_id: message.messageId,
      message_send_user_id: message.messageSendUserId,
      message_send_date: message.messageSendTime,
      message_channel_id: message.messageChannelId,
      message_edit_flag: true,
      message_message: inputText
    };
    const json = JSON.stringify(putMessage);

    const headers = {
      "Content-Type": "application/json"
    };

    axios.put(url, json, { headers: headers }).then((res) => {
      console.log(res);
      setInputText("");
    });
    // .catch((err) => {
    //   console.log(err);
    // });
  };

  const onKeyPress = (event: KeyboardEvent<HTMLInputElement>) => {
    onClick();
  };

  const onClickModal = () => {
    setShowModal(false);
  };
  const onClickDelete = () => {
    let url = URLs.deleteMessage;
    url += message.messageId;
    axios.delete(url).then((res) => {
      console.log(res);
      setShowModal(false);
    });
  };

  if (showModal) {
    return (
      <Modal onClick={onClickModal}>
        <InputButton
          value={inputText}
          placeholder=""
          onChange={onChageInput}
          onKeyPress={onKeyPress}
          onClick={onClickMessage}
        >
          edit
        </InputButton>
        <SecondaryButton onClick={onClickDelete}>Delete</SecondaryButton>
      </Modal>
    );
  } else {
    return (
      <>
        <BaseContainer>
          <ShowUserName userId={message.messageSendUserId} />
          <ChatMessageButton onClick={onClickMessage}>
            {message.messageMessage}
          </ChatMessageButton>
        </BaseContainer>
        <BaseContainer>
          <p>{message.messageEditFlag ? "edit" : null}</p>
          <p>{message.messageSendTime}</p>
        </BaseContainer>
      </>
    );
  }
};
