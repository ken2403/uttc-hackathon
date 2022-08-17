import { FC, ChangeEvent, KeyboardEvent } from "react";
import styled from "styled-components";
import { PrimaryButton } from "../atoms/button/PrimaryButton";
import { BaseInput } from "../atoms/input/BaseInput";

const SInput = styled(BaseInput)`
  margin-right: 10px;
`;

type Props = {
  value: string | number;
  placeholder: string;
  onChange: (event: ChangeEvent<HTMLInputElement>) => void;
  onKeyPress: (event: KeyboardEvent<HTMLInputElement>) => void;
  onClick: () => void;
  children: React.ReactNode;
};

export const InputButton: FC<Props> = (props) => {
  const {
    value = "",
    placeholder = "",
    onChange = (e: ChangeEvent<HTMLInputElement>) => {},
    onKeyPress = (e: KeyboardEvent<HTMLInputElement>) => {},
    onClick = () => {},
    children
  } = props;
  return (
    <>
      <SInput
        value={value}
        placeholder={placeholder}
        onChange={onChange}
        onKeyPress={onKeyPress}
      />
      <PrimaryButton onClick={onClick}>{children}</PrimaryButton>
    </>
  );
};
