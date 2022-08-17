import { FC, KeyboardEvent, ChangeEvent } from "react";
import { InputButton } from "../molecules/InputButton";

type Props = {
  onKeyPress: (event: KeyboardEvent<HTMLInputElement>) => void;
  onClick: () => void;
  onChange: (event: ChangeEvent<HTMLInputElement>) => void;
  value: string;
};

export const SendMessagArea: FC<Props> = (props) => {
  const { onKeyPress, onClick, onChange, value } = props;

  return (
    <InputButton
      value={value}
      placeholder={"message"}
      onChange={onChange}
      onClick={onClick}
      onKeyPress={onKeyPress}
    >
      Send
    </InputButton>
  );
};
