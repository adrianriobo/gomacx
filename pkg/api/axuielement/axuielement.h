#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>

const char* HasAXUIElementChildren(CFTypeRef axuielement);
CFArrayRef GetAXUIElementChildren(CFTypeRef axuielement);
CFTypeRef GetChild(CFArrayRef children, CFIndex index);
const char* GetTitleAttribute(CFTypeRef axuielement);
const char* GetValueAttribute(CFTypeRef axuielement);
const char* GetRoleAttribute(CFTypeRef axuielement);