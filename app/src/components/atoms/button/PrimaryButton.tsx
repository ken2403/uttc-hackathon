import { FC } from "react";
import styled from "styled-components";
import { BaseButton } from "./BaseButton";

const SButton = styled(BaseButton)`
  background-color: #400100;
`;

type Props = {
  onClick: () => void;
  children: React.ReactNode;
};

export const PrimaryButton: FC<Props> = (props) => {
  const { onClick = () => {}, children } = props;
  return <SButton onClick={onClick}>{children}</SButton>;
};
