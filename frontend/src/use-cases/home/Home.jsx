import React from "react";

import {
  DigitLayout,
  DigitText,
  DigitDesign,
  DigitButton
} from "@cthit/react-digit-components";

const Home = ({}) => (
  <DigitLayout.Column>
    <DigitLayout.Center>
      <DigitDesign.Card absWidth="300px">
        <DigitDesign.CardHeader>
          <DigitDesign.CardTitle text="Avtal för digITs tjänster" />
        </DigitDesign.CardHeader>
        <DigitDesign.CardBody>
          <DigitLayout.Row fill justifyContent="space-around">
            <p>Vänster</p>
            <p>Höger</p>
          </DigitLayout.Row>
        </DigitDesign.CardBody>
        <DigitDesign.CardButtons reverseDirection>
          <DigitButton primary raised text="Hej" />
        </DigitDesign.CardButtons>
      </DigitDesign.Card>
    </DigitLayout.Center>
  </DigitLayout.Column>
);

export default Home;
