import { buildModule } from '@nomicfoundation/hardhat-ignition/modules';

export default buildModule('TodoModule', (m) => {
  const todo = m.contract('Todo');

  m.call(todo, 'addTask', ['Learn Solidity']);

  return { todo };
});
