import React, { useState } from "react";
import { Button, Modal, Header } from "semantic-ui-react";

function GameAlertModal({ msg, buttonText, handleButtonClick }) {
  const [open, setOpen] = useState(true);
  const onButtonClick = () => {
    handleButtonClick()
    setOpen(false)
  };
  // Simply pop-up message UI
  return (
    <>
      <Modal open={open} size="small">
        <Header>{msg}</Header>
        <Modal.Actions>
          <Button onClick={ onButtonClick }>{buttonText}</Button>
        </Modal.Actions>
      </Modal>
    </>
  );
}

export default GameAlertModal;