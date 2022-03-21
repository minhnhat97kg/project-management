import React from 'react';
import './App.less';
import { RouterNavigation } from './routes';
export declare type RangeValue<DateType> = [EventValue<DateType>, EventValue<DateType>] | null;
export declare type EventValue<DateType> = DateType | null;

// interface IOk<DateType> {
//   onOkay: (value: RangeValue<DateType>) => void;
// }

function App(): React.ReactElement {
  return (
    <div className="App">
      <RouterNavigation />
    </div>
  );
}

export default App;
