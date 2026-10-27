// Preserve the ordinary PureScript implementation on JavaScript. The native
// backend uses the same interface to select its shared construction plan.
export const recordNilImpl = _left => _support => right => _object => _proxy => right({});
export const recordConsImpl = _reflect => _step => _tail => fallback => _isRight => _rightValue => _right => _left => _support => _nothing => _just => fallback;

// Schema tags are a native-backend optimisation. JavaScript keeps the ordinary
// decoder and ignores the tag; behaviour is unchanged.
export const typedInt = _support => decoder => decoder;
export const typedJson = _support => decoder => decoder;
export const typedNumber = _support => decoder => decoder;
export const typedString = _support => decoder => decoder;
export const typedBoolean = _support => decoder => decoder;
export const typedMaybe = _support => _inner => decoder => decoder;
export const typedArray = _support => _inner => decoder => decoder;
export const typedRecord = _support => _row => decoder => decoder;
export const typedObject = _inner => decoder => _right => decoder;
export const typedFieldId = _inner => impl => impl;
export const typedFieldMaybe = _inner => impl => impl;
export const fieldStep = _field => step => step;
export const decodeFieldFast = decoder => json => decoder(json);
export const getFieldImpl = fallback => fallback;
export const getFieldOptionalImpl = fallback => fallback;
export const getFieldOptionalNullableImpl = fallback => fallback;
export const borrowObject = decoder => json => decoder(json);
export const schemaDecoderABI1 = 1;
