import { Header } from '@/components/layout/header';
import { Main } from '@/components/layout/main';
import { ProfileDropdown } from '@/components/profile-dropdown';
import { Search } from '@/components/search';
import { ThemeSwitch } from '@/components/theme-switch';
import { columns } from './components/algorithms-columns';
import { AlgorithmsTable } from './components/algorithms-table';
import { UsersDialogs } from './components/users-dialogs';
import AlgorithmsContext from './context/algorithms-context';
import { algorithms } from './data/algorithms';

export default function Algorithms() {
  // Parse user list

  return (
    <>
      <Header fixed>
        <Search />
        <div className="ml-auto flex items-center space-x-4">
          <ThemeSwitch />
          <ProfileDropdown />
        </div>
      </Header>
      <AlgorithmsContext>
        <Main>
          <div className="mb-2 flex flex-wrap items-center justify-between space-y-2">
            <div>
              <h2 className="text-2xl font-bold tracking-tight">Algorithms</h2>
              <p className="text-muted-foreground">Manage, configure and deploy your algorithms here.</p>
            </div>
          </div>
          <div className="-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-y-0 lg:space-x-12">
            <AlgorithmsTable data={algorithms} columns={columns} />
          </div>
        </Main>
        <UsersDialogs />
      </AlgorithmsContext>
    </>
  );
}
