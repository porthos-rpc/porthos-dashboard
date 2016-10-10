import React from 'react';
import {render} from 'react-dom';
import Menu from './components/menu'

class App extends React.Component {
  render () {
    return (
        <div>
            <Menu />
        </div>
    );
  }
}

render(<App/>, document.getElementById('app'));
