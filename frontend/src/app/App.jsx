import React, { Component } from "react";

import { Switch, Route } from "react-router";

import {
  DigitProviders,
  DigitHeader,
  DigitNavLink
} from "@cthit/react-digit-components";

import Home from "../use-cases/home";

import Agreement from "../use-cases/agreement";

const DemoAgreement = `
Det här är ett dummy användaravtal. Det betyder inprincip ingenting. Det är för att kunna se i avtal.chalmers.it hur ett användaravtal kan se ut. Det finns alltså ingenting att acceptera här alltså.

Jag hade kunnat lura er och skrivit om ni använder avtal.chalmers.it så behöver ni betala en massa pengar, men det tänker jag inte göra. Just nu försöker jag bara komma på saker att skriva för att få texten att bli längre, nu kommer en punktlista.

- Glass kan vara godare än godis, men choklad är godast.
- Göteborg är den finaste staden av dem andra.
- Vad kan man skriva i en punktlista egentligen?
- En sista liten punkt.

Here we go again. Det här är ett dummy användaravtal. Det betyder inprincip ingenting. Det är för att kunna se i avtal.chalmers.it hur ett användaravtal kan se ut. Det finns alltså ingenting att acceptera här alltså.

Jag hade kunnat lura er och skrivit om ni använder avtal.chalmers.it så behöver ni betala en massa pengar, men det tänker jag inte göra. Just nu försöker jag bara komma på saker att skriva för att få texten att bli längre, nu kommer en punktlista.

- Glass kan vara godare än godis, men choklad är godast.
- Göteborg är den finaste staden av dem andra.
- Vad kan man skriva i en punktlista egentligen?
- En sista liten punkt.
`;

class App extends Component {
  render() {
    return (
      <DigitProviders>
        <DigitHeader
          title="Avtal.chalmers.it"
          renderDrawer={closeDrawer => (
            <nav>
              <DigitNavLink onClick={closeDrawer} link="/" text="Hem" />
              <DigitNavLink
                onClick={closeDrawer}
                link="/accept"
                text="Acceptera avtal"
              />
              <DigitNavLink
                onClick={closeDrawer}
                link="/it-account"
                text="IT kontotjänster"
              />
              <DigitNavLink
                onClick={closeDrawer}
                link="/hubbit"
                text="hubbIT"
              />
              <DigitNavLink
                onClick={closeDrawer}
                link="/bookit"
                text="BookIT"
              />
            </nav>
          )}
          renderMain={() => (
            <Switch>
              <Route
                exact
                route="/it-account"
                render={() => (
                  <Agreement
                    title="Demo användaravtal"
                    agreement={DemoAgreement}
                  />
                )}
              />
              <Route exact route="/" component={Home} />
              {/* <Route exact route="/accept" /> */}
              {/* <Route exact route="/hubbit" /> */}
              {/* <Route exact route="/bookit" /> */}
            </Switch>
          )}
        />
      </DigitProviders>
    );
  }
}

export default App;
