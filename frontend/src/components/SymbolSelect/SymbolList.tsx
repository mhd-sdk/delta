import { ContainedList, ContainedListItem, Search } from '@carbon/react';
import { css } from '@emotion/css';
import { debounce } from 'lodash';
import { useEffect, useMemo, useState } from 'react';
import { GetSymbols } from '../../../wailsjs/go/main/App';

interface Props {
  onSelect: (symbol: string) => void;
}

export const SymbolList = ({ onSelect }: Props): JSX.Element => {
  const [symbols, setSymbols] = useState<string[]>([]);
  const [loading, setLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState<string[]>([]);

  const handleChange = useMemo(
    () =>
      debounce((event) => {
        setSearchTerm(event.target.value);
        setLoading(false);
      }, 300),
    []
  );

  useEffect(() => {
    // Filtrage optimisé : le rendu des résultats est recalculé seulement si nécessaire
    setSearchResults(symbols.filter((listItem) => listItem.toLowerCase().includes(searchTerm.toLowerCase())));
  }, [symbols, searchTerm]);

  useEffect(() => {
    const fetchSymbols = async () => {
      setLoading(true);
      const res = await GetSymbols();
      setSymbols(res.map(({ symbol }) => symbol));
      setSearchResults(res.map(({ symbol }) => symbol));
      setLoading(false);
    };
    fetchSymbols();
  }, []);

  const renderList = () => {
    if (loading) {
      return <div>Loading...</div>;
    }
    if (searchTerm.length === 0) {
      return <></>;
    }
    if (searchResults.length === 0) {
      return <div>No results found</div>;
    }
    return searchResults.slice(0, 50).map((symbol, key) => (
      <ContainedListItem className={styles.listItem} key={key} onClick={() => onSelect(symbol)}>
        {symbol}
      </ContainedListItem>
    ));
  };

  return (
    <div className={styles.layout}>
      <div className={styles.body}>
        <ContainedList
          kind="disclosed"
          label={
            <Search
              placeholder=""
              labelText=""
              defaultValue={searchTerm}
              onChange={(e) => {
                setLoading(true);
                handleChange(e);
              }}
              className={styles.search}
              closeButtonLabelText="Clear search input"
              size="sm"
            />
          }
          size="sm"
        >
          {renderList()}
        </ContainedList>
      </div>
      {/* <div className={styles.actions}>
        <Button size="sm" kind="danger--ghost" onClick={onBack} disabled={disabled}>
          Cancel
        </Button>
        <Button size="sm" disabled={disabled}>
          Apply
        </Button>
      </div> */}
    </div>
  );
};

const styles = {
  search: css`
    input {
      text-transform: uppercase !important;
    }
  `,
  listItem: css`
    cursor: pointer;
  `,
  body: css`
    flex: 1;
    overflow-y: auto;
  `,
  layout: css`
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 150px;
    max-height: 300px;
    .cds--contained-list__header {
      padding: 0;
    }
  `,
  actions: css`
    margin-left: auto;
  `,
  moreResults: css`
    margin-top: 1rem;
    font-size: 0.875rem;
    color: gray;
  `,
};
