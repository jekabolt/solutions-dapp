import { useState } from 'react';

import styles from 'styles/test.module.scss';

const Test = () => {
  const [sum, setSum] = useState(0);

  return (
    <div className={styles.container}>
      <h1>Test page</h1>
      <h3>admin panel for minting nfts</h3>
      <h5>sum: {sum}</h5>
      <button onClick={() => setSum((val) => ++val)}>+</button>
      <button onClick={() => setSum((val) => --val)}>-</button>
    </div>
  );
};

export default Test;
