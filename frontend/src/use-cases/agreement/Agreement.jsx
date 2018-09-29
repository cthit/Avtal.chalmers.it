import React from "react";

import {
  DigitLayout,
  DigitText,
  DigitMarkdown
} from "@cthit/react-digit-components";

const Agreement = ({ title, agreement }) => (
  <DigitLayout.Center>
    <DigitLayout.Size maxWidth="700px">
      <DigitLayout.Column fill padding={"8px"}>
        <DigitText.Display text={title} />
        <DigitMarkdown markdownSource={agreement} />
      </DigitLayout.Column>
    </DigitLayout.Size>
  </DigitLayout.Center>
);

export default Agreement;
