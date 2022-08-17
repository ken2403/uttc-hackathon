import { FC } from "react";
import styled from "styled-components";

const ModalContainer = styled.div`
  position: fixed;
  top: 50;
  left: 50;
  width: 70%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.8);
  /* display: flex; */
  vertical-align: ;
  justify-content: center;
  z-index: 0;
  padding: 1em;
  background: #fff;
`;

type Props = {
  onClick: () => void;
  children: React.ReactNode;
};

export const Modal: FC<Props> = (props) => {
  const { onClick, children } = props;
  return (
    <ModalContainer>
      {children}
      <br />

      <button onClick={onClick}>close</button>
    </ModalContainer>
  );
};
