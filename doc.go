/*
Package keypack packs values into composite keys, using a byte encoding that
preserves the sort order of the original values.

Many of the encoding formats are the same as those used in HBase's OrderedBytes
class, which were in turn modeled after encoding formats defined by SQLite4 and
Orderly.

Coming soon: This package includes codecs for time.Time, bool, and unsigned integers.
*/
package keypack
