import { FC } from "react";
import { BaseContainer } from "../atoms/container/BaseContainer";

type Props = {
  userName: string;
};

export const NoUserAlert: FC<Props> = (props) => {
  const { userName } = props;
  return (
    <BaseContainer>
      <h3>{userName}という名前のユーザーは存在しません</h3>
      <h3>正しいユーザー名を入力してください</h3>
    </BaseContainer>
  );
};
