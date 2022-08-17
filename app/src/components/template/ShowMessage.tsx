import { FC } from "react";

import { BaseContainer } from "../atoms/container/BaseContainer";
import { ShowOneChat } from "../organism/ShowOneChat";
import { Message } from "../../types//message";

type Props = {
  messages: Message[];
  userId: string;
};

export const ShowMessage: FC<Props> = (props) => {
  const { messages, userId } = props;

  console.log(messages);
  return (
    <BaseContainer>
      {messages.map((message) => (
        <ShowOneChat
          key={message.messageId}
          message={message}
          userId={userId}
        />
      ))}
    </BaseContainer>
  );
};
