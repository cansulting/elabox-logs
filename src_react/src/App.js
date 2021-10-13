import logo from './logo.svg';
import './App.css';
import { ChakraProvider} from '@chakra-ui/react'
import Logs from "./components/logTable"

function App() {
  return (
    <ChakraProvider>
      <div className="App">
        {/* <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </header> */}
        <Logs />
      </div>
    </ChakraProvider>
  );
}

export default App;
